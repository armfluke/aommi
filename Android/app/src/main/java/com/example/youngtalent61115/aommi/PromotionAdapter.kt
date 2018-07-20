package com.example.youngtalent61115.aommi

import android.content.Context
import android.support.v7.widget.RecyclerView
import android.util.Log
import android.view.LayoutInflater
import android.view.MotionEvent
import android.view.View
import android.view.ViewGroup
import com.staytuned.mo.tngptutorial.networking.PromotionDataResponse
import kotlinx.android.synthetic.main.promotion_list_item.view.*

class PromotionAdapter(val list: /*ArrayList<String>*/List<PromotionDataResponse>, val context: Context) : RecyclerView.Adapter<ViewHolder>() {
    override fun onCreateViewHolder(p0: ViewGroup, p1: Int): ViewHolder {

        return ViewHolder(LayoutInflater.from(context).inflate(R.layout.promotion_list_item, p0, false))
    }

    override fun getItemCount(): Int {
        return list.size
    }

    override fun onBindViewHolder(p0: ViewHolder, p1: Int) {
        p0?.tvPromotionName?.text = list[p1].promotionName
        p0?.tvPromotionUsePoint?.text = list[p1].point.toString()
        Log.d("armfluke", "!!!!"+list[p1]!!.promotionName)
    }

}

class ViewHolder (view: View) : RecyclerView.ViewHolder(view){
    val tvPromotionName = view.tvPromotionName
    val tvPromotionUsePoint = view.tvPromotionUsePoint

}