package com.staytuned.mo.tngptutorial.networking

import retrofit2.Call
import retrofit2.http.FieldMap
import retrofit2.http.GET
import retrofit2.http.POST
import retrofit2.http.Query

interface AommiApi {
    @GET("/promotion")
    fun getPromotion(/*@Query("after") after: String,
               @Query("limit") limit: String*/)
            : Call<PromotionResponse>

    @POST("/promotion/use")
    fun usePromotion(@FieldMap params: MutableMap<String, String>)
            : Call<PromotionResponse>
}