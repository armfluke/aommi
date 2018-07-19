package com.staytuned.mo.tngptutorial.networking

import retrofit2.Call
import retrofit2.http.FieldMap
import retrofit2.http.GET
import retrofit2.http.POST
import retrofit2.http.Query

interface RedditApi {
    @GET("/top.json")
    fun getTop(@Query("after") after: String,
               @Query("limit") limit: String)
            : Call<RedditNewsResponse>

    @POST("/top.json")
    fun saveData(@FieldMap params: MutableMap<String, String>)
            : Call<RedditNewsResponse>
}