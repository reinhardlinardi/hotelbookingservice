package booking;

public class Guest {
    public String name;
    public int phone;
    public String idCard;
    public String email;

    Guest(){

    }

    Guest(String name_,int phone_,String idCard_, String email_){
        this.name = name_;
        this.phone = phone_;
        this.idCard = idCard_;
        this.email = email_;
    }
}
