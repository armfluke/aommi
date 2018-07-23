package com.example.youngtalent61115.aommi.networking

import android.os.Parcel
import android.os.Parcelable

data class Account(
        val accountID: String,
        val accountName: String,
        val pointBalance: Int
): Parcelable {
    constructor(parcel: Parcel) : this(
            parcel.readString(),
            parcel.readString(),
            parcel.readInt()) {
    }

    override fun writeToParcel(parcel: Parcel, flags: Int) {
        parcel.writeString(accountID)
        parcel.writeString(accountName)
        parcel.writeInt(pointBalance)
    }

    override fun describeContents(): Int {
        return 0
    }

    companion object CREATOR : Parcelable.Creator<Account> {
        override fun createFromParcel(parcel: Parcel): Account {
            return Account(parcel)
        }

        override fun newArray(size: Int): Array<Account?> {
            return arrayOfNulls(size)
        }
    }
}