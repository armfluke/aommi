package com.staytuned.mo.tngptutorial.networking

class AccountResponse(val data: List<AccountDataResponse>) {
    class AccountDataResponse(
            val accountID: String,
            val accountName: String,
            val pointBalance: Int
    )
}