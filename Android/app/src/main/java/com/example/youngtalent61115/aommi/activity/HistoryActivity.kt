package com.example.youngtalent61115.aommi.activity

import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.support.v7.widget.LinearLayoutManager
import android.support.v7.widget.RecyclerView
import android.util.Log
import android.view.View
import android.widget.LinearLayout
import com.beust.klaxon.Klaxon
import com.example.youngtalent61115.aommi.GlobalVariable
import com.example.youngtalent61115.aommi.HistoryViewAdapter
import com.example.youngtalent61115.aommi.PromotionViewAdapter
import com.example.youngtalent61115.aommi.R
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.History
import com.example.youngtalent61115.aommi.networking.Promotion
import com.github.kittinunf.fuel.android.extension.responseJson
import com.github.kittinunf.fuel.httpPost
import com.github.kittinunf.result.Result
import kotlinx.android.synthetic.main.activity_history.*
import kotlinx.android.synthetic.main.activity_main.*

class HistoryActivity : AppCompatActivity(){
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_history)

        showProgressCircle()

        val account = intent.getParcelableExtra<Account>("account")

        "${GlobalVariable.baseUrl}/account/history".httpPost().body("{\"accountID\":\"${account.accountID}\"}").responseString() { request, response, result ->
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
                    try {
                        val history = ArrayList(Klaxon().parseArray<History>(data))
                        createRecyclerView(account, history)
                    }catch (e: Exception){

                    }

                    hideProgressCircle()
                }
            }
        }
    }

    private fun createRecyclerView(account: Account, history: ArrayList<History>) {
        val rv = findViewById<RecyclerView>(R.id.rcvHistoryList)
        rv.layoutManager = LinearLayoutManager(this, LinearLayout.VERTICAL, false)

        var adapter = HistoryViewAdapter(this, account, history)
        rv.adapter = adapter
    }

    private fun hideProgressCircle() {
        rcvHistoryList.visibility = View.VISIBLE
        progressBar.visibility = View.GONE
    }

    private fun showProgressCircle() {
        rcvHistoryList.visibility = View.GONE
        progressBar.visibility = View.VISIBLE
    }
}
