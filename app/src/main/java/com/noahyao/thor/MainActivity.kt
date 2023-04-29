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
        val port = "8888"
        var uri = "ping"

        Server.listenAtBackground(port)


        // Load WebView
        val webView = findViewById<WebView>(R.id.webView)
        webView.settings.javaScriptEnabled = true
        webView.loadUrl("$host:$port/$uri")
    }
}