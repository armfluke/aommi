package com.example.youngtalent61115.aommi.activity

import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import com.example.youngtalent61115.aommi.R
import kotlinx.android.synthetic.main.activity_main.*
import kotlinx.android.synthetic.main.activity_reward.*
import java.util.*
import android.content.DialogInterface
import android.support.v7.app.AlertDialog
import android.widget.Toast
import com.example.youngtalent61115.aommi.MainActivity



class RewardActivity : AppCompatActivity(){
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_reward)
        clickRedeem()

    }

    private fun clickRedeem() {
        btn_redeem.setOnClickListener {
            //to detail
            val builder = AlertDialog.Builder(this@RewardActivity)
            val poinForRedeem = decreasePoint(1000,100)
            var messageForSucces = ""
            if(poinForRedeem >= 0){
                messageForSucces = "คุณได้ทำการใช้คอยน์เพื่อแลกบัตรดูหนังในเครือเมเจอร์ " + poinForRedeem.toString()+" coin"
            }else messageForSucces = "ขออภัย จำนวนพอยต์ของคุณไม่พอ"
            builder.setMessage(messageForSucces)
            builder.setPositiveButton("ยืนยัน", DialogInterface.OnClickListener { dialog, id -> finish(

            )
                val intent = Intent(applicationContext, ShowCodeRedeemActivity::class.java)
                var redeemCode = generateCode(8)
                intent.putExtra("RedeemCode",redeemCode)
                startActivity(intent)
            })


            builder.setNegativeButton("ยกเลิก") { dialog, id -> dialog.cancel()
                //dialog.dismiss();
            }
            builder.setTitle("Redeem")

            builder.show()
        }
    }
    fun decreasePoint(currentBalance: Int, decreasePoint: Int): Int {
        return currentBalance - decreasePoint
    }
    fun generateCode(positionCount: Int): String{
        val chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
        var passWord = ""
        for (i in 0..positionCount) {
            passWord += chars[Math.floor(Math.random() * chars.length).toInt()]
        }
        return passWord
    }
}