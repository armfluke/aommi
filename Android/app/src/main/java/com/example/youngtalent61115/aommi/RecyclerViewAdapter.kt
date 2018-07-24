package com.example.youngtalent61115.aommi

import android.content.Intent
import android.support.v4.content.ContextCompat.startActivity
import android.support.v7.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import android.widget.Toast
import com.bumptech.glide.Glide
import com.example.youngtalent61115.aommi.activity.RewardActivity
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.Promotion

class RecyclerViewAdapter(val mainActivity: MainActivity, val account: Account, val promotionList: ArrayList<Promotion>): RecyclerView.Adapter<RecyclerViewAdapter.ViewHolder>() {

    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        holder?.promotionName?.text = promotionList[position].promotionName
        //holder?.promotionImage?.text = promotionList[position].image
        holder?.description?.text = promotionList[position].condition
        holder?.coin?.text = promotionList[position].point.toString()

        Glide.with(mainActivity).load(promotionList[position].image).into(holder?.promotionImage)

        holder?.redeemButton?.setOnClickListener{
            //Toast.makeText(mainActivity, "Clicked!!!", Toast.LENGTH_LONG).show()

            val intent = Intent(mainActivity, RewardActivity::class.java)
            intent.putExtra("account", account)
            intent.putExtra("promotion", promotionList[position])
            startActivity(mainActivity ,intent, null)
        }
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        val v = LayoutInflater.from(parent?.context).inflate(R.layout.item_list, parent, false)
        return ViewHolder(v);
    }

    override fun getItemCount(): Int {
        return promotionList.size
    }

    class ViewHolder(itemView: View): RecyclerView.ViewHolder(itemView){
        val promotionName = itemView.findViewById<TextView>(R.id.promotionName)
        //val promotionImage = itemView.findViewById<TextView>(R.id.promotionImage)
        val description = itemView.findViewById<TextView>(R.id.promotionDescription)
        val coin = itemView.findViewById<TextView>(R.id.coin)
        val redeemButton = itemView.findViewById<TextView>(R.id.redeemButton)
        val promotionImage = itemView.findViewById<ImageView>(R.id.promotionImage)
    }

}