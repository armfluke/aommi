package com.example.youngtalent61115.aommi.networking

import android.os.Parcel
import android.os.Parcelable

data class Promotion(
        val promotionID: Int,
        val promotionName: String,
        val reward: String,
        val condition: String,
        val point: Int,
        val image: String,
        val description: String,
        val limitUse: Int
): Parcelable {
    constructor(parcel: Parcel) : this(
            parcel.readInt(),
            parcel.readString(),
            parcel.readString(),
            parcel.readString(),
            parcel.readInt(),
            parcel.readString(),
            parcel.readString(),
            parcel.readInt()) {
    }

    override fun writeToParcel(parcel: Parcel, flags: Int) {
        parcel.writeInt(promotionID)
        parcel.writeString(promotionName)
        parcel.writeString(reward)
        parcel.writeString(condition)
        parcel.writeInt(point)
        parcel.writeString(image)
        parcel.writeString(description)
        parcel.writeInt(limitUse)
    }

    override fun describeContents(): Int {
        return 0
    }

    companion object CREATOR : Parcelable.Creator<Promotion> {
        override fun createFromParcel(parcel: Parcel): Promotion {
            return Promotion(parcel)
        }

        override fun newArray(size: Int): Array<Promotion?> {
            return arrayOfNulls(size)
        }
    }
}