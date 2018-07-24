package com.example.youngtalent61115.aommi.activity

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.util.Log
import com.example.youngtalent61115.aommi.GlobalVariable
import com.example.youngtalent61115.aommi.R
import com.example.youngtalent61115.aommi.networking.Account
import com.github.kittinunf.fuel.httpPost
import com.github.kittinunf.result.Result
import kotlinx.android.synthetic.main.activity_show_point.*
import java.util.*
import java.text.SimpleDateFormat


class ShowPointActivity : AppCompatActivity(){
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_show_point)

        val account = intent.getParcelableExtra<Account>("account")
        //val qrCode = intent.getStringExtra("qrCode")
        val qrPoint = intent.getIntExtra("qrPoint", 0)

        txt_show_receive_point.text = qrPoint.toString()
        txt_show_balance_point.text = (account.pointBalance + qrPoint).toString()

        clickAccecpt()
        setDateQRCode()
        //requestToUpdatePoint(account, qrCode)
    }
    private fun clickAccecpt(){

        btn_accept_earn_point.setOnClickListener {
            this.finish()
        }
    }
    private fun setEarnPoint(){
        txt_show_receive_point.text = "50"
    }
    private fun setBalancePoint(){
        txt_show_balance_point.text = "300"
    }

    private fun setDateQRCode(){
        txt_code_date_exp.text = getCurrentDateTime()
    }
    private fun getCurrentDateTime(): String{
        val cal = Calendar.getInstance()
        val date = cal.time
        val dateFormat = SimpleDateFormat("yyyy-MM-dd HH:mm:ss")
        val formattedDate = dateFormat.format(date)
        return  formattedDate
    }

//    private fun requestToUpdatePoint(account: Account, qrCode: String){
//        val body = "{\"qrCode\":\"${qrCode}\",\"accountID\":\"${account.accountID}\"}"
//        "${GlobalVariable.baseUrl}/qr".httpPost().body(body).responseString{ request, response, result ->
//            //do something with response
//            when (result) {
//                is Result.Failure -> {
//                    val ex = result.getException()
//                    Log.d("armfluke", "Fail!!!")
//                    Log.d("armfluke", ex.toString())
//                }
//                is Result.Success -> {
//                    val data = result.get()
//
//                    Log.d("armfluke", "Success!!!")
//                    Log.d("armfluke", data)
//                }
//            }
//        }
//    }
}
