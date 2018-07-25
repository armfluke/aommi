package com.example.youngtalent61115.aommi.activity

import android.os.Bundle
import android.os.PersistableBundle
import android.support.v7.app.AppCompatActivity
import android.util.Log
import com.beust.klaxon.Klaxon
import com.example.youngtalent61115.aommi.GlobalVariable
import com.example.youngtalent61115.aommi.R
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.History
import com.github.kittinunf.fuel.httpPost
import com.github.kittinunf.result.Result

class HistoryActivity : AppCompatActivity(){
    override fun onCreate(savedInstanceState: Bundle?, persistentState: PersistableBundle?) {
        super.onCreate(savedInstanceState, persistentState)
        setContentView(R.layout.activity_history)
        Log.d("xxx","xxx")

        val account = intent.getParcelableExtra<Account>("account")

        "${GlobalVariable.baseUrl}/account/history".httpPost().body("\"accountID\":\"${account.accountID}\"").responseString(){ request, response, result ->
            //do something with response
            when (result) {
                is Result.Failure -> {
                    val ex = result.getException()
                    Log.d("armfluke", "Fail!!!")
                    Log.d("armfluke", ex.toString())
                }
                is Result.Success -> {
                    val data = result.get()
                    Log.d("armfluke2", "Success!!!")
                    Log.d("armfluke2", data)
                    val history = ArrayList(Klaxon().parseArray<History>(data))


                }
            }
        }
    }
}