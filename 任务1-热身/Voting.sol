// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {

    string[] public candidateList;

    // 一个mapping来存储候选人的得票数
    mapping(string => uint256) public candidateVoteCounts;
    // 一个vote函数，允许用户投票给某个候选人
    function vote(string memory _candidate) public {
        candidateVoteCounts[_candidate] += 1;
    }
    // 一个getVotes函数，返回某个候选人的得票数
    function getVotes(string memory _candidate) public view returns (uint256) {
        return candidateVoteCounts[_candidate];
    }
    // 一个resetVotes函数，重置所有候选人的得票数
    function resetVotes(string[] memory _candidates) public {
        for (uint i =0; i < _candidates.length; i++) {
           candidateVoteCounts[_candidates[i]] = 0;
        }
    }
}