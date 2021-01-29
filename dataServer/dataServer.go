package dataServer

import (
	"context"
	"log"

	"github.com/berrydata/BerryMiner/common"
	"github.com/berrydata/BerryMiner/config"
	"github.com/berrydata/BerryMiner/db"
	"github.com/berrydata/BerryMiner/rest"
	"github.com/berrydata/BerryMiner/rpc"
	"github.com/berrydata/BerryMiner/tracker"
	"github.com/berrydata/BerryMiner/util"
)

//DataServer holds refs to primary stack of utilities for data retrieval and serving
type DataServer struct {
	server       *rest.Server
	DB           db.DB
	runner       *tracker.Runner
	ethClient    rpc.ETHClient
	exitCh       chan int
	runnerExitCh chan int
	Stopped      bool
	readyChannel chan bool
	log          *util.Logger
}

//CreateServer creates a data server stack and kicks off all go routines to start retrieving and serving data
func CreateServer(ctx context.Context) (*DataServer, error) {
	cfg := config.GetConfig()

	DB := ctx.Value(common.DBContextKey).(db.DB)
	client := ctx.Value(common.ClientContextKey).(rpc.ETHClient)
	run, err := tracker.NewRunner(client, DB)
	if err != nil {
		log.Fatal(err)
	}
	srv, err := rest.Create(ctx, cfg.ServerHost, cfg.ServerPort)

	//make sure channel buffer size 1 since there is no guarantee that anyone
	//would be listening to the channel
	ready := make(chan bool, 1)
	return &DataServer{server: srv,
		DB:           DB,
		runner:       run,
		ethClient:    client,
		exitCh:       nil,
		Stopped:      true,
		runnerExitCh: nil,
		readyChannel: ready,
		log:          util.NewLogger("dataServer", "DataServer")}, nil
}

//Start the data server and all underlying resources
func (ds *DataServer) Start(ctx context.Context, exitCh chan int) error {
	ds.exitCh = exitCh
	ds.runnerExitCh = make(chan int)
	ds.Stopped = false
	err := ds.runner.Start(ctx, ds.runnerExitCh)
	if err != nil {
		return err
	}

	ds.server.Start()
	go func() {
		<-ds.runner.Ready()
		ds.log.Info("Runner signaled it is ready")
		ds.readyChannel <- true
		ds.log.Info("DataServer ready for use")
		<-ds.exitCh
		ds.log.Info("DataServer received signal to stop")
		ds.stop()
	}()
	return nil
}

//Ready provides notification channel that data from trackers is ready for use
func (ds *DataServer) Ready() chan bool {
	return ds.readyChannel
}

func (ds *DataServer) stop() error {
	//stop tracker run loop
	ds.runnerExitCh <- 1

	//stop REST erver
	ds.server.Stop()

	//stop the DB
	ds.DB.Close()

	//stop the eth RPC client
	ds.ethClient.Close()

	//all done
	ds.Stopped = true
	return nil
}
