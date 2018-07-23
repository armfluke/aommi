package com.example.youngtalent61115.aommi

import android.content.Context
import android.support.v7.widget.RecyclerView
import android.util.Log
import android.view.LayoutInflater
import android.view.MotionEvent
import android.view.View
import android.view.ViewGroup
import android.widget.Toast
import kotlinx.android.synthetic.main.promotion_list_item.view.*

class PromotionAdapter(val list: ArrayList<String>, val context: Context) : RecyclerView.Adapter<ViewHolder>() {
    override fun onCreateViewHolder(p0: ViewGroup, p1: Int): ViewHolder {

        return ViewHolder(LayoutInflater.from(context).inflate(R.layout.promotion_list_item, p0, false))
    }

    override fun getItemCount(): Int {
        return list.size
    }

    override fun onBindViewHolder(p0: ViewHolder, p1: Int) {
        p0?.tvPromotionName?.text = list[p1]

        p0.itemView.setOnClickListener{
            Log.d("beer","name: ${p0.tvPromotionName.text}")
        }
    }

}

class ViewHolder (itemView: View) : RecyclerView.ViewHolder(itemView){
    val tvPromotionName = itemView.tvPromotionName

}