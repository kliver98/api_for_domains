package com.kliver.domains.api.model;

public class Server {

    private String address;
    private String ssl_grade;
    private String country;
    private String owner;

    public Server(){}

    public Server(String address, String sslGrade, String country, String owner) {
        this.address = address;
        this.ssl_grade = sslGrade;
        this.country = country;
        this.owner = owner;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public String getSsl_grade() {
        return ssl_grade;
    }

    public void setSsl_grade(String sslGrade) {
        this.ssl_grade = sslGrade;
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
