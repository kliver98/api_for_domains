package com.kliver.domains.api.model;

public class History {

    private String []items;

    public History(){}

    public History(String[] items) {
        this.items = items;
    }

    public String[] getItems() {
        return items;
    }

    public void setItems(String[] items) {
        this.items = items;
    }

}
