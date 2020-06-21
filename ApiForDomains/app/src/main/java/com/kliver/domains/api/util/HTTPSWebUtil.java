package com.kliver.domains.api.util;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.SocketTimeoutException;
import java.net.URL;

public class HTTPSWebUtil {

    private OnResponseListener listener;


    public HTTPSWebUtil() {}

    public void setListener(OnResponseListener listener) {
        this.listener = listener;
    }

    public interface OnResponseListener {
        void onResponse(int callbackID, String response);
    }

    public void GETrequest(int callbackID, String url) {
        try {
            URL page = new URL(url);
            HttpURLConnection connection = (HttpURLConnection) page.openConnection();
            connection.setConnectTimeout(Constants.TIMEOUT_HTTP_URL_CONNECTION);
            InputStream is = connection.getInputStream();
            ByteArrayOutputStream baos = new ByteArrayOutputStream();
            byte[] buffer = new byte[4096];
            int bytesRead;
            while ((bytesRead = is.read(buffer)) != -1) {
                baos.write(buffer, 0, bytesRead);
            }
            is.close();
            baos.close();
            connection.disconnect();
            String response = new String(baos.toByteArray(), "UTF-8");
            if (listener != null) listener.onResponse(callbackID, response);
        }catch (SocketTimeoutException ex) {
            if (listener != null) listener.onResponse(callbackID, Constants.CONNECTION_TIMEOUT_RESPONSE);
        }catch (IOException ex){
            ex.printStackTrace();
        }
    }

}