package com.example.youngtalent61115.aommi

import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.support.v7.widget.LinearLayoutManager
import android.util.Log
import android.view.View
import com.staytuned.mo.tngptutorial.networking.AommiApi
import com.staytuned.mo.tngptutorial.networking.PromotionResponse
import com.staytuned.mo.tngptutorial.networking.RedditNewsResponse
import com.staytuned.mo.tngptutorial.networking.RestAPI
import com.example.youngtalent61115.aommi.activity.RewardActivity
import kotlinx.android.synthetic.main.activity_main.*
import retrofit2.Call
import retrofit2.Response
import retrofit2.Callback

class MainActivity : AppCompatActivity() {

    private val promotion: ArrayList<String> = ArrayList()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        addPromotions()
        setBalancePoint()
        setRecyclerView()
        loadService()

        clickPromotion()



    }

    private fun clickPromotion() {
        relPromotion1.setOnClickListener {
            //to detail
            val intent = Intent(applicationContext, RewardActivity::class.java)
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

    private fun addPromotions() {
        promotion.add("หลวงพี่แจ๊ส 5G")
        promotion.add("หลวงพี่แจ๊ส 6G")
    }

    private fun loadService() {
        RestAPI().create().getPromotion().enqueue(object : Callback<PromotionResponse> {
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

        })
    }
}
