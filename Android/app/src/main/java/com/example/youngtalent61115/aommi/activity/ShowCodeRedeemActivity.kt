package com.example.youngtalent61115.aommi.activity

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import com.example.youngtalent61115.aommi.R
import kotlinx.android.synthetic.main.activity_show_code_redeem.*

class ShowCodeRedeemActivity  : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_show_code_redeem)
        var redeemCode: String = intent.getStringExtra("RedeemCode")
        txt_redeem_code.setText(redeemCode)

    }
}