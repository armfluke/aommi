package com.example.youngtalent61115.aommi

import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.support.v7.widget.LinearLayoutManager
import android.util.Log
import kotlinx.android.synthetic.main.activity_main.*

class MainActivity : AppCompatActivity() {

    private val promotion: ArrayList<String> = ArrayList()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        addPromotions()
        setBalancePoint()
        setRecyclerView()

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

}
