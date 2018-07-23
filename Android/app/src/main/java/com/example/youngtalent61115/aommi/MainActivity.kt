package com.example.youngtalent61115.aommi

import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.support.v7.widget.LinearLayoutManager
import android.util.Log
import android.view.View
import com.beust.klaxon.Klaxon
import com.example.youngtalent61115.aommi.activity.RewardActivity
import com.github.kittinunf.fuel.httpGet
import com.github.kittinunf.result.Result
import com.staytuned.mo.tngptutorial.networking.*
import kotlinx.android.synthetic.main.activity_main.*
import retrofit2.Call
import retrofit2.Response
import retrofit2.Callback

class MainActivity : AppCompatActivity() {

    private var promotion: List<PromotionDataResponse> = ArrayList()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        getAllPromotion()

        //addPromotions()
        setBalancePoint()
        //setRecyclerView()

        clickPromotion()
    }

    private fun getAllPromotion(){
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
                    Log.d("armfluke", data)
                    val pro = Klaxon().parseArray<PromotionDataResponse>(data)
                    Log.d("armfluke", pro!![0].promotionName)
                    //promotion = pro.toMutableList()
                    //Log.d("armfluke", promotion!![0].promotionName)
                    tvPromotionName.text = pro!![0].promotionName
                    tvPromotionUsePoint.text = pro!![0].point.toString()
                    //setRecyclerView(promotion)
                    //rcvPromotionList.adapter?.notifyDataSetChanged()

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

    private fun setRecyclerView() {
        rcvPromotionList.layoutManager = LinearLayoutManager(this)
        rcvPromotionList.adapter = PromotionAdapter(promotion, this)
    }

    /*private fun addPromotions() {
        promotion.add("หลวงพี่แจ๊ส 5G")
        promotion.add("หลวงพี่แจ๊ส 6G")
    }*/

}
