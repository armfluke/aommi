package com.example.youngtalent61115.aommi.activity

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.util.Log
import com.beust.klaxon.Klaxon
import com.example.youngtalent61115.aommi.R
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.Promotion
import com.github.kittinunf.fuel.httpPost
import com.github.kittinunf.result.Result
import kotlinx.android.synthetic.main.activity_show_code_redeem.*

class ShowCodeRedeemActivity  : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_show_code_redeem)
        var redeemCode: String = intent.getStringExtra("redeemCode")
        txt_redeem_code.setText(redeemCode)

        val promotion = intent.getParcelableExtra<Promotion>("promotion")
        val account = intent.getParcelableExtra<Account>("account")

        usePromotion(account, promotion, redeemCode)
    }

    fun usePromotion(account: Account, promotion: Promotion, redeemCode: String){
        val bodyPromotionUse = "{\"accountID\":\"" + account.accountID + "\",promotionID\":" + promotion.promotionID + ",\"rewardCode\":\"" + redeemCode + "\",\"pointBalance\":" + (account.pointBalance - promotion.point).toString() + "}"
        //Log.d("armfluke", bodyPromotionUse)
        "http://10.0.2.2:3000/promotion/use".httpPost().body(bodyPromotionUse).responseString{ request, response, result ->
            //do something with response
            when (result) {
                is Result.Failure -> {
                    val ex = result.getException()
                    Log.d("armfluke", "Fail!!!")
                    Log.d("armfluke", ex.toString())
                }
                is Result.Success -> {
                    val data = result.get()

                    Log.d("armfluke", "Success!!!")
                    Log.d("armfluke", data)
                }
            }
        }

        val bodyAccount = "{\"accountID\":\"" + account.accountID + "\",\"pointBalance\":" + (account.pointBalance - promotion.point).toString() + "}"
        "http://10.0.2.2:3000/point/update".httpPost().body(bodyAccount).responseString{ request, response, result ->
            //do something with response
            when (result) {
                is Result.Failure -> {
                    val ex = result.getException()
                    Log.d("armfluke", "Fail!!!")
                    Log.d("armfluke", ex.toString())
                }
                is Result.Success -> {
                    val data = result.get()

                    Log.d("armfluke", "Success!!!")
                    Log.d("armfluke", data)
                }
            }
        }
    }
}