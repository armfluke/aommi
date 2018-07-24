package com.example.youngtalent61115.aommi.activity

import android.accounts.Account
import android.content.Intent
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.widget.Toast
import com.budiyev.android.codescanner.AutoFocusMode
import com.budiyev.android.codescanner.CodeScanner
import com.budiyev.android.codescanner.DecodeCallback
import com.budiyev.android.codescanner.ErrorCallback
import com.budiyev.android.codescanner.ScanMode
import com.example.youngtalent61115.aommi.R
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
                //Open ShowPoint
                val intent = Intent(applicationContext, ShowPointActivity::class.java)
                intent.putExtra("account", account)
                intent.putExtra("qrCode", it.text)
                startActivity(intent)
                this.finish()
                Toast.makeText(this, "Scan result: ${it.text}", Toast.LENGTH_LONG).show()
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
}
