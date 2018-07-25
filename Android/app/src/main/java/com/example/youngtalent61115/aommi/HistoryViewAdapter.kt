package com.example.youngtalent61115.aommi

import android.content.Intent
import android.support.v4.content.ContextCompat.startActivity
import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import com.bumptech.glide.Glide
import com.example.youngtalent61115.aommi.activity.HistoryActivity
import com.example.youngtalent61115.aommi.activity.RewardActivity
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.History
import com.example.youngtalent61115.aommi.networking.Promotion

class HistoryViewAdapter(val historyActivity: HistoryActivity, val account: Account, val historyList: ArrayList<History>): RecyclerView.Adapter<HistoryViewAdapter.ViewHolder>() {

    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        holder?.promotionName?.text = historyList[position].reward
        holder?.point?.text = historyList[position].point.toString()
        holder?.date?.text = historyList[position].date
        holder?.code?.text = "Redeem : ${historyList[position].code}"

        Glide.with(historyActivity).load(historyList[position].image).into(holder?.promotionImage)

    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        val v = LayoutInflater.from(parent?.context).inflate(R.layout.item_history_list, parent, false)
        return ViewHolder(v);
    }

    override fun getItemCount(): Int {
        return historyList.size
    }

    class ViewHolder(itemView: View): RecyclerView.ViewHolder(itemView){
        val promotionName = itemView.findViewById<TextView>(R.id.tvPromotionName)
        val point = itemView.findViewById<TextView>(R.id.tvRedeemfor)
        val date = itemView.findViewById<TextView>(R.id.tvUsedDate)
        val code = itemView.findViewById<TextView>(R.id.tvCode)
        val promotionImage = itemView.findViewById<ImageView>(R.id.promotionImage)
    }

}