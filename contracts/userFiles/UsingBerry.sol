// pragma solidity ^0.5.0;

// import '../Berry.sol';
// import '../BerryMaster.sol';
// import './UserContract.sol';
// /**
// * @title UsingBerry
// * This contracts creates for easy integration to the Berry Berry System
// */
// contract UsingBerry{
// 	UserContract berryUserContract;
// 	address public owner;
	
// 	event OwnershipTransferred(address _previousOwner,address _newOwner);


//     constructor(address _user)public {
//     	berryUserContract = UserContract(_user);
//     	owner = msg.sender;
//     }

// 	function getCurrentValue(uint _requestId) public view returns(bool ifRetrieve, uint value, uint _timestampRetrieved) {
// 		BerryMaster _berry = BerryMaster(berryUserContract.berryStorageAddress());
// 		uint _count = _berry.getNewValueCountbyRequestId(_requestId) ;
// 		if(_count > 0){
// 				_timestampRetrieved = _berry.getTimestampbyRequestIDandIndex(_requestId,_count -1);//will this work with a zero index? (or insta hit?)
// 				return(true,_berry.retrieveData(_requestId,_timestampRetrieved),_timestampRetrieved);
//         }
//         return(false,0,0);
// 	}

// 	//How can we make this one more efficient?
// 	function getFirstVerifiedDataAfter(uint _requestId, uint _timestamp) public returns(bool,uint,uint _timestampRetrieved) {
// 		BerryMaster _berry = BerryMaster(berryUserContract.berryStorageAddress());
// 		uint _count = _berry.getNewValueCountbyRequestId(_requestId);
// 		if(_count > 0){
// 				for(uint i = _count;i > 0;i--){
// 					if(_berry.getTimestampbyRequestIDandIndex(_requestId,i-1) > _timestamp && _berry.getTimestampbyRequestIDandIndex(_requestId,i-1) < block.timestamp - 86400){
// 						_timestampRetrieved = _berry.getTimestampbyRequestIDandIndex(_requestId,i-1);//will this work with a zero index? (or insta hit?)
// 					}
// 				}
// 				if(_timestampRetrieved > 0){
// 					return(true,_berry.retrieveData(_requestId,_timestampRetrieved),_timestampRetrieved);
// 				}
//         }
//         return(false,0,0);
// 	}
	
// 	event Print(string _s,uint _num);

// 	function getAnyDataAfter(uint _requestId, uint _timestamp) public  returns(bool _ifRetrieve, uint _value, uint _timestampRetrieved){
// 		BerryMaster _berry = BerryMaster(berryUserContract.berryStorageAddress());
// 		uint _count = _berry.getNewValueCountbyRequestId(_requestId) ;
// 		if(_count > 0){
// 				emit Print("count",_count);
// 				for(uint i = _count;i > 0;i--){
// 					emit Print('tester',_berry.getTimestampbyRequestIDandIndex(_requestId,i-1));
// 					emit Print('actual', _timestamp);

// 					if(_berry.getTimestampbyRequestIDandIndex(_requestId,i-1) >= _timestamp){
// 						_timestampRetrieved = _berry.getTimestampbyRequestIDandIndex(_requestId,i-1);//will this work with a zero index? (or insta hit?)
// 						emit Print("_timestampRetrieved",_timestampRetrieved);
// 						emit Print("value",_berry.retrieveData(_requestId,_timestampRetrieved));
// 					}
// 				}
// 				if(_timestampRetrieved > 0){
// 					return(true,_berry.retrieveData(_requestId,_timestampRetrieved),_timestampRetrieved);
// 				}
//         }
//         return(false,0,0);
// 	}

// 	function requestData(string calldata _request,string calldata _symbol,uint _granularity, uint _tip) external{
// 		Berry _berry = Berry(berryUserContract.berryStorageAddress());
// 		if(_tip > 0){
// 			require(_berry.transferFrom(msg.sender,address(this),_tip));
// 		}
// 		_berry.requestData(_request,_symbol,_granularity,_tip);
// 	}

// 	function requestDataWithEther(string calldata _request,string calldata _symbol,uint _granularity, uint _tip) payable external{
// 		berryUserContract.requestDataWithEther.value(msg.value)(_request,_symbol,_granularity,_tip);
// 	}

// 	function addTip(uint _requestId, uint _tip) public {
// 		Berry _berry = Berry(berryUserContract.berryStorageAddress());
// 		require(_berry.transferFrom(msg.sender,address(this),_tip));
// 		_berry.addTip(_requestId,_tip);
// 	}

// 	function addTipWithEther(uint _requestId, uint _tip) public payable {
// 		UserContract(berryUserContract).addTipWithEther.value(msg.value)(_requestId,_tip);
// 	}

// 	function setUserContract(address _userContract) public {
// 		require(msg.sender == owner);//who should this be?
// 		berryUserContract = UserContract(_userContract);
// 	}

// 	function transferOwnership(address payable _newOwner) external {
//             require(msg.sender == owner);
//             emit OwnershipTransferred(owner, _newOwner);
//             owner = _newOwner;
//     }
// }


