package com.staytuned.mo.tngptutorial.networking

import retrofit2.Retrofit
import retrofit2.converter.moshi.MoshiConverterFactory

class RestAPI {
    fun create(): RedditApi {


        val retrofit = Retrofit.Builder()
                .addConverterFactory(MoshiConverterFactory.create())
                .baseUrl("https://192.168.43.225:3000")
                .build()

        return retrofit.create(RedditApi::class.java)
    }
}