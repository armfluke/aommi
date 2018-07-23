package com.example.youngtalent61115.aommi

import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.support.v7.widget.LinearLayoutManager
import android.support.v7.widget.RecyclerView
import android.util.Log
import android.view.View
import android.widget.LinearLayout
import com.beust.klaxon.Klaxon
import com.example.youngtalent61115.aommi.activity.RewardActivity
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.Promotion
import com.example.youngtalent61115.aommi.activity.ScanQRActivity
import com.github.kittinunf.fuel.httpGet
import com.github.kittinunf.result.Result
import com.staytuned.mo.tngptutorial.networking.PromotionDataResponse
import kotlinx.android.synthetic.main.activity_main.*

class MainActivity : AppCompatActivity() {

    //private val promotion: ArrayList<String> = ArrayList()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        //setBalancePoint()
        getAccount()
        clickScanQR()
    }

    fun createRecyclerView(account: Account, promotion: ArrayList<Promotion>){
        val rv = findViewById<RecyclerView>(R.id.rcvPromotionList)
        rv.layoutManager = LinearLayoutManager(this, LinearLayout.VERTICAL, false)

        var adapter = RecyclerViewAdapter(this, account, promotion)
        rv.adapter = adapter
    }

    private fun getAccount(){
        "http://10.0.2.2:3000/account".httpGet().responseString(){ request, response, result ->
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
                    tvBalancePoint.text = account[0].pointBalance.toString()
                    account_name.text = account[0].accountName

                    getAllPromotion(account[0])
                }
            }
        }
        //AccountID, AccountName, PointBalance
    }

    private fun clickScanQR() {
        btnScanQR.setOnClickListener {
            val intent = Intent(applicationContext, ScanQRActivity::class.java)
            startActivity(intent)
        }
    }

    private fun getAllPromotion(account: Account){
        //an extension over string (support GET, PUT, POST, DELETE with httpGet(), httpPut(), httpPost(), httpDelete())
        "http://10.0.2.2:3000/promotion".httpGet().responseString { request, response, result ->
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

    private fun clickPromotion() {
        relPromotion1.setOnClickListener {
            //to detail
            val intent = Intent(applicationContext, RewardActivity::class.java)
            intent.putExtra("PROMOTION_ID","")
            startActivity(intent)
        }
    }

    private fun setBalancePoint() {
        tvBalancePoint.text = "300"
    }

//    private fun setRecyclerView() {
//        rcvPromotionList.layoutManager = LinearLayoutManager(this)
//        rcvPromotionList.adapter = PromotionAdapter(promotion, this)
//    }
//
//    private fun addPromotions() {
//        promotion.add("Avengers 1")
//        promotion.add("Avengers 2")
//        promotion.add("Avengers 3")
//        promotion.add("Avengers 4")
//        promotion.add("Avengers 5")
//        promotion.add("Avengers 6")
//        promotion.add("Avengers 7")
//        promotion.add("Avengers 8")
//
//    }

    private fun loadService() {
        /*RestAPI().create().getPromotion(object: Callback<List<PromotionDataResponse>> {
            override fun success(call: Call<PromotionResponse>?, response: Response<PromotionResponse>?) {
                if (response!!.isSuccessful) {
                    Log.d("armfluke", "Successful!!!!!!")
                    val promo = response.body()
                    Log.d("armfluke", promo.toString())
                }
            }
        })*/

        /*override fun onFailure(call: Call<PromotionResponse>?, t: Throwable?) {
            Log.d("armfluke", "Fail!!!!!!")
            Log.d("armfluke", t.toString())
        })*//*.enqueue(object : Callback<List<PromotionDataResponse>> {
            override fun onResponse(call: Call<PromotionResponse>?, response: Response<PromotionResponse>?) {
                if (response!!.isSuccessful) {
                    Log.d("armfluke", "Successful!!!!!!")
                    val promo = response.body()
                    Log.d("armfluke", promo.toString())
                }
            }

            override fun onFailure(call: Call<PromotionResponse>?, t: Throwable?) {
                Log.d("armfluke", "Fail!!!!!!")
                Log.d("armfluke", t.toString())
            }

        })*/
    }
}
