package com.example.youngtalent61115.aommi.activity

import android.content.Intent
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.support.v7.app.AlertDialog
import android.util.Log
import android.widget.Toast
import com.beust.klaxon.Klaxon
import com.budiyev.android.codescanner.AutoFocusMode
import com.budiyev.android.codescanner.CodeScanner
import com.budiyev.android.codescanner.DecodeCallback
import com.budiyev.android.codescanner.ErrorCallback
import com.budiyev.android.codescanner.ScanMode
import com.example.youngtalent61115.aommi.GlobalVariable
import com.example.youngtalent61115.aommi.R
import com.example.youngtalent61115.aommi.networking.Account
import com.example.youngtalent61115.aommi.networking.QRResponse
import com.github.kittinunf.fuel.httpPost
import com.github.kittinunf.result.Result
import kotlinx.android.synthetic.main.activity_scan_qr.*

class ScanQRActivity : AppCompatActivity() {

    private lateinit var codeScanner: CodeScanner

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_scan_qr)

        val account = intent.getParcelableExtra<Account>("account")

        setScannerProperties()
        callbackFromScanner(account)

        clickScannerToContinual()
    }

    private fun clickScannerToContinual() {
        scanner_view.setOnClickListener {
            codeScanner.startPreview()
        }
    }

    private fun callbackFromScanner(account: Account) {
        codeScanner.decodeCallback = DecodeCallback {
            runOnUiThread {

                requestToCheckAndUpdatePoint(account, it.text)

                //Open ShowPoint
                /*val intent = Intent(applicationContext, ShowPointActivity::class.java)
                intent.putExtra("account", account)
                intent.putExtra("qrCode", it.text)
                startActivity(intent)
                this.finish()
                Toast.makeText(this, "Scan result: ${it.text}", Toast.LENGTH_LONG).show()*/
            }
        }
        codeScanner.errorCallback = ErrorCallback {
            runOnUiThread {
                Toast.makeText(this, "Camera initialization error: ${it.message}",
                        Toast.LENGTH_LONG).show()
            }
        }
    }

    private fun setScannerProperties() {
        codeScanner = CodeScanner(this, scanner_view)

        codeScanner.camera = CodeScanner.CAMERA_BACK
        codeScanner.formats = CodeScanner.ALL_FORMATS
        codeScanner.autoFocusMode = AutoFocusMode.SAFE
        codeScanner.scanMode = ScanMode.SINGLE
        codeScanner.isAutoFocusEnabled = true
        codeScanner.isFlashEnabled = false
    }

    override fun onResume() {
        super.onResume()
        codeScanner.startPreview()
    }

    override fun onPause() {
        codeScanner.releaseResources()
        super.onPause()
    }

    private fun requestToCheckAndUpdatePoint(account: Account, qrCode: String){
        val body = "{\"qrCode\":\"${qrCode}\",\"accountID\":\"${account.accountID}\"}"
        "${GlobalVariable.baseUrl}/qr".httpPost().body(body).responseString{ request, response, result ->
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

                    val qr = Klaxon().parse<QRResponse>(data)

                    if(qr?.qrPoint == 0){
                        val builder = AlertDialog.Builder(this@ScanQRActivity)

                        builder.setNegativeButton("ยกเลิก") { dialog, id ->
                            dialog.cancel()
                            codeScanner.startPreview()
                        }

                        builder.setMessage("ขออภัย ไม่พบ QR Code ในระบบ")
                        builder.setTitle("Scan QR Code")

                        builder.show()
                    }else{
                        val intent = Intent(applicationContext, ShowPointActivity::class.java)
                        intent.putExtra("account", account)
                        //intent.putExtra("qrCode", qrCode)
                        intent.putExtra("qrPoint", qr?.qrPoint)
                        startActivity(intent)
                        this.finish()
                    }
                }
            }
        }
    }
}
