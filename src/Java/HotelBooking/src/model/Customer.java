package model;

public class Customer {

    /* Properties */
    int id;
    String name;
    String identity;
    String email;

    // Constructors
    public Customer() {}

    public Customer(int id, String name, String identity, String email){
        this.id = id;
        this.name = name;
        this.identity = identity;
        this.email = email;
    }

    /* Methods */

    /* Getters & Setters */

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getIdentity() {
        return identity;
    }

    public void setIdentity(String identity) {
        this.identity = identity;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}
