package com.kliver.domains.api.control;

import android.util.Log;
import android.view.View;
import android.widget.ArrayAdapter;
import com.google.gson.Gson;
import com.kliver.domains.api.R;
import com.kliver.domains.api.model.History;
import com.kliver.domains.api.util.Constants;
import com.kliver.domains.api.util.HTTPSWebUtil;
import com.kliver.domains.api.view.HistoryActivity;

import java.util.ArrayList;

public class HistoryController implements View.OnClickListener, HTTPSWebUtil.OnResponseListener {

    private HistoryActivity activity;
    private HTTPSWebUtil util;


    public HistoryController(HistoryActivity activity) {
        this.activity = activity;
        util = new HTTPSWebUtil();
        activity.getBackArrowIV().setOnClickListener(this);
        initializeHistory();
    }

    public void initializeHistory() {
        String url = Constants.URL_HISTORY;
        Thread t = new Thread(
                () -> {
                    util.GETrequest(Constants.HISTORY_CALLBACK, url);
                }
        );
        t.start();
        try {
            t.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    @Override
    public void onClick(View v) {
        switch (v.getId()) {
            case R.id.backArrowIV:
                activity.onBackPressed();
                break;
        }
    }

    @Override
    public void onResponse(int callbackID, String response) {
        System.out.println(">>>>Hola"+callbackID+":"+response);
        switch (callbackID) {
            case Constants.HISTORY_CALLBACK:
                Gson gson = new Gson();
                History history = gson.fromJson(response, History.class);
                activity.runOnUiThread(
                        () -> {
                            for (String item:history.getItems()) {
                                activity.getItems().add(item);
                            }
                            activity.getAdapter().notifyDataSetChanged();
                            Log.e(">>>",activity.getItems().size()+"");
                        }
                );

                break;
        }
    }
}
