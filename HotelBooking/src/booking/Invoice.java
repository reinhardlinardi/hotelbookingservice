package booking;


public class Invoice {
    public String name;
    public String date;
    public String paymentStatus;
    public String paymentMethod;
    public String idInvoice;
    public String idBooking;
    public int price;
    public boolean breakfast;

    public Invoice(){

    }

    public Invoice(String name_,
            String date_,
            String paymentStatus_,
            String paymentMethod_,
            String idInvoice_,
            String idBooking_,
            int price_,
            boolean breakfast_){
        this.name = name_;
        this.date = date_;
        this.paymentStatus = paymentStatus_;
        this.paymentMethod = paymentMethod_;
        this.idInvoice = idInvoice_;
        this.idBooking = idBooking_;
        this.price = price_;
        this.breakfast = breakfast_;

    }
}
