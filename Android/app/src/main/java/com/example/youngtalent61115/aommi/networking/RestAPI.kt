package com.staytuned.mo.tngptutorial.networking

import retrofit2.Retrofit
import retrofit2.converter.moshi.MoshiConverterFactory

class RestAPI {
    fun create(): RedditApi {

        val retrofit = Retrofit.Builder()
                .addConverterFactory(MoshiConverterFactory.create())
                .baseUrl("https://www.reddit.com")
                .build()

        return retrofit.create(RedditApi::class.java)
    }
}