package com.example.youngtalent61115.aommi.activity

import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import com.example.youngtalent61115.aommi.R
import kotlinx.android.synthetic.main.activity_reward.*
import android.content.DialogInterface
import android.graphics.Color
import android.support.v7.app.AlertDialog
import android.util.Log
import com.bumptech.glide.Glide
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.Promotion


class RewardActivity : AppCompatActivity(){
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_reward)

        val promotion = intent.getParcelableExtra<Promotion>("promotion")
        val account = intent.getParcelableExtra<Account>("account")

        name_promotion.text = promotion.promotionName
        piont_for_redeem.text = promotion.point.toString() + " coins"
        condition_promotion.text = promotion.condition
        detail_promotion.text = promotion.description

        setLimit(promotion)
        Glide.with(this).load(promotion.image).into(promotionImage)
        promotionBackground.setBackgroundColor(Color.WHITE)

        clickRedeem(account, promotion)

    }

    private fun setLimit(promotion: Promotion) {
        if (promotion.limitUse.toString() == "0") {
            limit_use.text = "-"
        } else {
            limit_use.text = promotion.limitUse.toString()
        }
    }

    private fun clickRedeem(account: Account, promotion: Promotion) {
        btn_redeem.setOnClickListener {
            //to detail
            val builder = AlertDialog.Builder(this@RewardActivity)
            val pointForRedeem = decreasePoint(account.pointBalance, promotion.point)
            var messageForSuccess = ""
            var redeemCode = ""
            if(pointForRedeem >= 0){
                messageForSuccess = "คุณได้ทำการใช้คอยน์สำหรับ " + promotion.promotionName+ " เป็นจำนวน " + promotion.point.toString()+ " coins คงเหลือ " + pointForRedeem.toString() + " coins"
                redeemCode = generateCode(8)
                Log.d("reward", redeemCode)

                builder.setPositiveButton("ยืนยัน", DialogInterface.OnClickListener { dialog, id ->
                    val intent = Intent(applicationContext, ShowCodeRedeemActivity::class.java)
                    intent.putExtra("redeemCode", redeemCode)
                    intent.putExtra("promotion", promotion)
                    intent.putExtra("account", account)
                    this.finish()
                    startActivity(intent)
                })

                builder.setNegativeButton("ยกเลิก") { dialog, id ->
                    dialog.cancel()
                }

            }else {
                messageForSuccess = "ขออภัย จำนวนคอยน์ของคุณไม่พอ"

                builder.setNegativeButton("ยืนยัน", DialogInterface.OnClickListener { dialog, id ->
                    dialog.cancel()
                })
            }
            builder.setMessage(messageForSuccess)
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