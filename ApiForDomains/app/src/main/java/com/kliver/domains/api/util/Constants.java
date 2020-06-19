package com.kliver.domains.api.util;

public class Constants {

    public static final int HOST_CALLBACK = 1;
    public static final int HISTORY_CALLBACK = 2;

    public static final String PORT = "8082";
    public static final String LOCALHOST = "http://192.168.1.5";
    public static final String API_BASE = "/api/v1/";
    public static final String GET_DOMAIN = API_BASE+"domain=:domain";
    public static final String GET_HISTORY = API_BASE+"history";

    public static final String URL_DOMAIN = LOCALHOST+":"+PORT+GET_DOMAIN;
    public static final String URL_HISTORY = LOCALHOST+":"+PORT+GET_HISTORY;
}