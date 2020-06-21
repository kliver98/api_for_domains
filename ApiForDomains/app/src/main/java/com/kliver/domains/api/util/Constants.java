package com.kliver.domains.api.util;

public class Constants {

    public static final int PING_CALLBACK = 0;
    public static final int DOMAIN_CALLBACK = 1;
    public static final int HISTORY_CALLBACK = 2;
    public static final int TIMEOUT_HTTP_URL_CONNECTION = 100; //100ms=0.1s
    public static final String CONNECTION_TIMEOUT_RESPONSE = "Connection Timeout - "+TIMEOUT_HTTP_URL_CONNECTION+" ms";

    public static final String PORT = "8082";
    public static final String LOCALHOST = "http://192.168.1.5";
    public static final String API_BASE = "/api/v1/";
    public static final String GET_DOMAIN = API_BASE+"domain=";
    public static final String GET_HISTORY = API_BASE+"history";
    public static final String URL_DOMAIN = LOCALHOST+":"+PORT+GET_DOMAIN;
    public static final String URL_HISTORY = LOCALHOST+":"+PORT+GET_HISTORY;

    public static final String DOMAIN_VAR_INTENT = "domain";
    public static final String GENERAL_ERROR = "Ha ocurrido un error: ";
    public static final String NO_DATA = "No hay datos ";

}