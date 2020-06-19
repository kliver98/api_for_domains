package com.kliver.domains.api.model;

public class Server {

    private String address;
    private String sslGrade;
    private String country;
    private String owner;

    public Server(){}

    public Server(String address, String sslGrade, String country, String owner) {
        this.address = address;
        this.sslGrade = sslGrade;
        this.country = country;
        this.owner = owner;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public String getSslGrade() {
        return sslGrade;
    }

    public void setSslGrade(String sslGrade) {
        this.sslGrade = sslGrade;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public String getOwner() {
        return owner;
    }

    public void setOwner(String owner) {
        this.owner = owner;
    }

}
