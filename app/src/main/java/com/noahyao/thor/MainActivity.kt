package com.noahyao.thor

import android.os.Bundle
import android.webkit.WebView
import androidx.appcompat.app.AppCompatActivity
import server.Server


class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)


        var host = "http://localhost"
        val port = "9000"
        var uri = "fakehtml/reports.html"
        requestPermissions(arrayOf("android.permission.WRITE_EXTERNAL_STORAGE"), 100)
        Server.listenAtBackground(port)

        // Load WebView
        val webView = findViewById<WebView>(R.id.webView)
        webView.settings.javaScriptEnabled = true
        webView.loadUrl("$host:$port/$uri")
    }

    override fun onRequestPermissionsResult(
        requestCode: Int, permissions: Array<out String>, grantResults: IntArray
    ) {
        super.onRequestPermissionsResult(requestCode, permissions, grantResults)
    }
}