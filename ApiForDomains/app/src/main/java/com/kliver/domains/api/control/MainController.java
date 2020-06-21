package com.kliver.domains.api.control;

import android.content.Context;
import android.util.Log;
import android.view.View;
import android.widget.Toast;

import com.kliver.domains.api.R;
import com.kliver.domains.api.util.Constants;
import com.kliver.domains.api.util.HTTPSWebUtil;
import com.kliver.domains.api.view.MainActivity;

public class MainController implements View.OnClickListener, HTTPSWebUtil.OnResponseListener {

    private MainActivity activity;
    private Boolean apiAlive;

    public MainController(MainActivity activity) {
        this.activity = activity;
        apiAlive = false;

        activity.getBtnDomain().setOnClickListener(this);
        activity.getBtnHistory().setOnClickListener(this);
    }

    @Override
    public void onClick(View v) {
        HTTPSWebUtil util = new HTTPSWebUtil();
        util.setListener(this);
        Thread t = new Thread(
                () -> {
                    util.GETrequest(Constants.PING_CALLBACK,Constants.URL_HISTORY);
                }
        );
        t.start();
        try {
            t.join();
        } catch (InterruptedException e) {
            Toast.makeText(activity,Constants.GENERAL_ERROR+"\n"+e.getMessage(),Toast.LENGTH_LONG).show();
        }
        switch(v.getId()) {
            case R.id.btnDomain:
                if (apiAlive)
                    activity.openDomainActivity();
                else
                    return;
                break;
            case R.id.btnHistory:
                if (apiAlive)
                    activity.openHistoryActivity();
                else
                    return;
                break;
        }
    }

    @Override
    public void onResponse(int callbackID, String response) {
        switch (callbackID) {
            case Constants.PING_CALLBACK:
                if (response.equals(Constants.CONNECTION_TIMEOUT_RESPONSE)) {
                    activity.runOnUiThread(
                            () -> {
                                Toast.makeText(activity,"Error: "+response,Toast.LENGTH_LONG).show();
                            }
                    );
                }
                else
                    apiAlive = true;
                break;
        }
    }
}
