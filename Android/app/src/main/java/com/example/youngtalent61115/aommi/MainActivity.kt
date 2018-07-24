package com.example.youngtalent61115.aommi

import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.support.v7.widget.LinearLayoutManager
import android.support.v7.widget.RecyclerView
import android.util.Log
import android.widget.LinearLayout
import com.beust.klaxon.Klaxon
import com.example.youngtalent61115.aommi.activity.RewardActivity
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.Promotion
import com.example.youngtalent61115.aommi.activity.ScanQRActivity
import com.github.kittinunf.fuel.httpGet
import com.github.kittinunf.result.Result
import kotlinx.android.synthetic.main.activity_main.*

object GlobalVariable {
    val baseUrl = "http://178.128.48.140:3001"
}

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        getAccount()
    }

    fun createRecyclerView(account: Account, promotion: ArrayList<Promotion>){
        val rv = findViewById<RecyclerView>(R.id.rcvPromotionList)
        rv.layoutManager = LinearLayoutManager(this, LinearLayout.VERTICAL, false) as RecyclerView.LayoutManager?

        var adapter = RecyclerViewAdapter(this, account, promotion)
        rv.adapter = adapter
    }

    private fun getAccount(){
        "${GlobalVariable.baseUrl}/account".httpGet().responseString(){ request, response, result ->
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
                    val account = ArrayList(Klaxon().parseArray<Account>(data))
                    setBalancePointAndAccountName(account)
                    getAllPromotion(account[0])
                    clickScanQR(account[0])
                }
            }
        }
        //AccountID, AccountName, PointBalance
    }

    private fun setBalancePointAndAccountName(account: ArrayList<Account>) {
        tvBalancePoint.text = account[0].pointBalance.toString()
        account_name.text = account[0].accountName
    }

    private fun clickScanQR(account: Account) {
        btnScanQR.setOnClickListener {
            val intent = Intent(applicationContext, ScanQRActivity::class.java)
            intent.putExtra("account", account)
            startActivity(intent)
        }
    }

    private fun getAllPromotion(account: Account){
        //an extension over string (support GET, PUT, POST, DELETE with httpGet(), httpPut(), httpPost(), httpDelete())
        "${GlobalVariable.baseUrl}/promotion".httpGet().responseString { request, response, result ->
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
                    createRecyclerView(account, ArrayList(Klaxon().parseArray<Promotion>(data)))

                }
            }
        }
    }


}
