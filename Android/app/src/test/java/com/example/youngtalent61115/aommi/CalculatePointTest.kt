package com.example.youngtalent61115.aommi.activity

import org.junit.Assert.*
import org.junit.Test

class CalculatePointTest {
    @Test
    fun decreasePointInputPointBalance1000AndPointReward100ShuldbePointBlance900Test() {
        val rewardCalculate = RewardActivity()
        val expectedResult = 900
        val actualResult = rewardCalculate.decreasePoint(1000, 100)
        assertEquals(expectedResult, actualResult)
    }
}

