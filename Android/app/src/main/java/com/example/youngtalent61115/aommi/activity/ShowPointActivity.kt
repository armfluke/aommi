package com.example.youngtalent61115.aommi.activity

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import com.example.youngtalent61115.aommi.R
import kotlinx.android.synthetic.main.activity_show_point.*
import java.util.*
import java.text.SimpleDateFormat


class ShowPointActivity : AppCompatActivity(){
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_show_point)
        clickAccecpt()
        setDateQRCode()

    }
    private fun clickAccecpt(){

        btn_accept_earn_point.setOnClickListener {
            this.finish()
        }
    }
    private fun setEarnPoint(){
        txt_show_receive_point.text = "50"
    }
    private fun setBalancePoint(){
        txt_show_balance_point.text = "300"
    }

    private fun setDateQRCode(){
        txt_code_date_exp.text = getCurrentDateTime()
    }
    private fun getCurrentDateTime(): String{
        val cal = Calendar.getInstance()
        val date = cal.time
        val dateFormat = SimpleDateFormat("yyyy-MM-dd HH:mm:ss")
        val formattedDate = dateFormat.format(date)
        return  formattedDate
    }
}
