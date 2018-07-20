package com.staytuned.mo.tngptutorial.networking

class PromotionResponse(val data: List<PromotionDataResponse>) {
    class PromotionDataResponse(
            val promotionID: Int,
            val promotionName: String,
            val reward: String,
            val condition: String,
            val point: Int,
            val image: String
    )
}