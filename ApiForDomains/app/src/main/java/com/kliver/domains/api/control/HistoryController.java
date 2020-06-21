package com.kliver.domains.api.control;

import android.view.View;
import com.google.gson.Gson;
import com.kliver.domains.api.R;
import com.kliver.domains.api.model.History;
import com.kliver.domains.api.util.Constants;
import com.kliver.domains.api.util.HTTPSWebUtil;
import com.kliver.domains.api.view.HistoryActivity;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class HistoryController implements View.OnClickListener, HTTPSWebUtil.OnResponseListener {

    private HistoryActivity activity;
    private HTTPSWebUtil util;


    public HistoryController(HistoryActivity activity) {
        this.activity = activity;
        util = new HTTPSWebUtil();
        util.setListener(this);
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
        switch (callbackID) {
            case Constants.HISTORY_CALLBACK:
                Gson gson = new Gson();
                List<String> tmpDATA = new ArrayList<String>();
                try {
                    tmpDATA = Arrays.asList(gson.fromJson(response, History.class).getItems());
                } catch (Exception e) {
                    if (response.equals("{}"))
                        tmpDATA = Arrays.asList(new String[]{Constants.NO_DATA});
                    else
                        tmpDATA = Arrays.asList(new String[]{"Error: "+response});
                }
                String[] data = new String[tmpDATA.size()];
                for (int i = 0; i < tmpDATA.size(); i++) {
                    data[i] = tmpDATA.get(i);
                }
                History history = new History();
                history.setItems(data);
                activity.runOnUiThread(
                        () -> {
                            for (String item:history.getItems()) {
                                activity.getItems().add(item);
                            }
                            activity.getAdapter().notifyDataSetChanged();
                        }
                );
                break;
        }
    }
}
