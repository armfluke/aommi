package com.example.youngtalent61115.aommi.activity

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import com.example.youngtalent61115.aommi.R

class RewardActivity : AppCompatActivity(){
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_reward)
    }
    fun decreasePoint(currentBalance: Int, decreasePoint: Int): Int {
        return currentBalance - decreasePoint
    }
}