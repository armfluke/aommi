package com.example.youngtalent61115.aommi.activity

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.util.Log
import com.example.youngtalent61115.aommi.GlobalVariable
import com.example.youngtalent61115.aommi.R
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.Promotion
import com.github.kittinunf.fuel.httpPost
import com.github.kittinunf.result.Result
import kotlinx.android.synthetic.main.activity_show_code_redeem.*
import java.text.SimpleDateFormat
import java.util.*

class ShowCodeRedeemActivity  : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_show_code_redeem)
        var redeemCode: String = intent.getStringExtra("redeemCode")
        txt_redeem_code.setText(redeemCode)

        val promotion = intent.getParcelableExtra<Promotion>("promotion")
        val account = intent.getParcelableExtra<Account>("account")

        usePromotion(account, promotion, redeemCode)
        setDateQRCode()
        clickCancel()
    }

    private fun clickCancel() {
        btn_cancel.setOnClickListener {
            this.finish()
        }
    }

    private fun setDateQRCode(){
        txt_date.text = getCurrentDateTime()
    }
    private fun getCurrentDateTime(): String{
        val cal = Calendar.getInstance()
        val date = cal.time
        val dateFormat = SimpleDateFormat("yyyy-MM-dd HH:mm")
        val formattedDate = dateFormat.format(date)
        return  formattedDate
    }

    fun usePromotion(account: Account, promotion: Promotion, redeemCode: String){
        val bodyPromotionUse = "{\"accountID\":\"${account.accountID}\",\"promotionID\":${promotion.promotionID},\"rewardCode\":\"${redeemCode}\"}"
        Log.d("ttt", bodyPromotionUse)
        "${GlobalVariable.baseUrl}/promotion/use".httpPost().body(bodyPromotionUse).responseString{ request, response, result ->
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
                    Log.d("ppp", request.toString())
                    Log.d("armfluke", data)
                }
            }
        }

        val bodyAccount = "{\"accountID\":\"" + account.accountID + "\",\"pointBalance\":" + (account.pointBalance - promotion.point).toString() + "}"
        "${GlobalVariable.baseUrl}/point/update".httpPost().body(bodyAccount).responseString{ request, response, result ->
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