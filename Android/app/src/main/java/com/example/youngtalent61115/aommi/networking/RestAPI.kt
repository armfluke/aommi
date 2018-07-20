package com.staytuned.mo.tngptutorial.networking

import retrofit2.Retrofit
import retrofit2.converter.moshi.MoshiConverterFactory

class RestAPI {
    fun create(): AommiApi {


        val retrofit = Retrofit.Builder()
                .addConverterFactory(MoshiConverterFactory.create())
                .baseUrl("http://10.0.2.2:3000")
                .build()

        return retrofit.create(AommiApi::class.java)
    }
}