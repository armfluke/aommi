package com.example.youngtalent61115.aommi.networking

import android.os.Parcel
import android.os.Parcelable

data class QRResponse(
        val balancePoint: Int,
        val qrPoint: Int
): Parcelable {
    constructor(parcel: Parcel) : this(
            parcel.readInt(),
            parcel.readInt()) {
    }

    override fun writeToParcel(parcel: Parcel, flags: Int) {
        parcel.writeInt(balancePoint)
        parcel.writeInt(qrPoint)
    }

    override fun describeContents(): Int {
        return 0
    }

    companion object CREATOR : Parcelable.Creator<QRResponse> {
        override fun createFromParcel(parcel: Parcel): QRResponse {
            return QRResponse(parcel)
        }

        override fun newArray(size: Int): Array<QRResponse?> {
            return arrayOfNulls(size)
        }
    }
}