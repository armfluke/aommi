package com.example.youngtalent61115.aommi.activity

import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import com.example.youngtalent61115.aommi.MainActivity
import com.example.youngtalent61115.aommi.R
import kotlinx.android.synthetic.main.activity_show_point.*
import java.time.LocalDateTime
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.time.format.FormatStyle
import java.time.temporal.TemporalQueries.zone
import java.util.*
import java.text.SimpleDateFormat


class ShowPointActivity : AppCompatActivity(){
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_show_point)
        clickAccecptAfterViewReceivePointAndBalancePoint()
        setExpireQRCode()


    }
    private fun clickAccecptAfterViewReceivePointAndBalancePoint(){
        val intent = Intent(applicationContext, MainActivity::class.java)
        startActivity(intent)
    }
    private fun setEarnPoint(){
        txt_show_receive_point.setText("50")
    }
    private fun setBalancePoint(){
        txt_show_balance_point.setText("300")
    }

    private fun setExpireQRCode(){
        txt_code_date_exp.setText(getCurrentDateTime())
    }
    private fun getCurrentDateTime(): String{
        val cal = Calendar.getInstance()
        val date = cal.time
        val dateFormat = SimpleDateFormat("HH:mm:ss")
        val formattedDate = dateFormat.format(date)
        return  formattedDate
    }
}
