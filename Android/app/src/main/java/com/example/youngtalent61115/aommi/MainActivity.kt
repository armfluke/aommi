package com.example.youngtalent61115.aommi

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.support.v7.widget.LinearLayoutManager
import kotlinx.android.synthetic.main.activity_main.*

class MainActivity : AppCompatActivity() {

    private val promotion: ArrayList<String> = ArrayList()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        addPromotions()
        setBalancePoint()
        setRecyclerView()



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
