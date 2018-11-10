package model;

import org.json.JSONObject;

import java.text.SimpleDateFormat;
import java.util.Date;

public class Invoice {

    /* Properties */

    int id;
    int roomId;
    int customerId;
    Date inDate;
    Date outDate;
    int price;
    boolean paid;
    boolean cancelled;

    // Constructors
    public Invoice(){}

    public Invoice(int id, int roomId, int customerId,Date inDate, Date outDate, int price, boolean paid, boolean cancelled) {
        this.id = id;
        this.roomId = roomId;
        this.customerId = customerId;
        this.inDate = inDate;
        this.outDate = outDate;
        this.price = price;
        this.paid = paid;
        this.cancelled = cancelled;
    }

    /* Methods */

    public void fillFromJSON(JSONObject json){
        this.roomId = json.getInt("room_id");
        this.customerId = json.getInt("customer_id");

        SimpleDateFormat formatter = new SimpleDateFormat("dd-MM-yyyy");
        try{
            this.inDate = formatter.parse(json.getString("in"));
            this.outDate = formatter.parse(json.getString("out"));
        }catch(Exception e){
            e.printStackTrace();
        }

        this.price = json.getInt("price");
        this.paid = json.getBoolean("paid");
        this.cancelled = json.getBoolean("cancelled");
    }

    public void printInvoice(){
        System.out.println("ID : "+this.id);
        System.out.println("Room ID : "+this.roomId);
        System.out.println("Customer ID : "+this.customerId);
        System.out.println("In Date : "+this.inDate);
        System.out.println("Out Date : "+this.outDate);
        System.out.println("Price : "+this.price);
        System.out.println("Paid : "+this.paid);
        System.out.println("Cancelled : "+this.cancelled);
    }

    /* Getters */
    public int getId() {
        return id;
    }

    public int getRoomId() {
        return roomId;
    }

    public int getCustomerId() {
        return customerId;
    }

    public Date getInDate() {
        return inDate;
    }

    public Date getOutDate() {
        return outDate;
    }

    public int getPrice() {
        return price;
    }

    public boolean getPaid() {
        return paid;
    }

    public boolean getCancelled() {
        return cancelled;
    }

    /* Setters */
    public void setId(int id) {
        this.id = id;
    }
}
