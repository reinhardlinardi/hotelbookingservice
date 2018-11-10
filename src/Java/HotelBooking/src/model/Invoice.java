package model;

import java.util.Date;

public class Invoice {

    /* Properties */

    int id;
    int roomId;
    int customerId;
    Date inDate;
    Date outDate;
    int price;
    int paid;
    int cancelled;

    /* Methods */

    // Constructors
    public Invoice(){}

    public Invoice(int id, int roomId, int customerId,Date inDate, Date outDate, int price, int paid, int cancelled) {
        this.id = id;
        this.roomId = roomId;
        this.customerId = customerId;
        this.inDate = inDate;
        this.outDate = outDate;
        this.price = price;
        this.paid = paid;
        this.cancelled = cancelled;
    }
}
