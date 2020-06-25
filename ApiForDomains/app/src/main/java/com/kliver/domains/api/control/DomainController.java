package com.kliver.domains.api.control;

import android.content.Context;
import android.view.View;
import android.view.inputmethod.InputMethodManager;
import android.widget.Toast;

import com.google.gson.Gson;
import com.kliver.domains.api.R;
import com.kliver.domains.api.model.Domain;
import com.kliver.domains.api.model.Server;
import com.kliver.domains.api.util.Constants;
import com.kliver.domains.api.util.HTTPSWebUtil;
import com.kliver.domains.api.view.DomainActivity;
import com.squareup.picasso.Picasso;

import java.util.List;

public class DomainController implements View.OnClickListener, HTTPSWebUtil.OnResponseListener {

    private DomainActivity activity;
    private String domain;
    private HTTPSWebUtil util;

    public DomainController(DomainActivity activity) {
        this.activity = activity;
        if (activity.getIntent().getExtras()!=null) {
            activity.getSearchBarET().setText(activity.getIntent().getExtras().getString(Constants.DOMAIN_VAR_INTENT));
        }
        util = new HTTPSWebUtil();
        util.setListener(this);
        initializeFileds();

        activity.getBackArrowIV().setOnClickListener(this);
        activity.getSearchImgIV().setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                if (view != null) {
                    //Hide the keyboard
                    InputMethodManager imm = (InputMethodManager)activity.getSystemService(Context.INPUT_METHOD_SERVICE);
                    imm.hideSoftInputFromWindow(view.getWindowToken(), 0);
                    //Search on API and put info
                    initializeFileds();
                }
            }
        });
    }

    private void initializeFileds() {
        domain = activity.getSearchBarET().getText().toString().trim();
        if (domain.trim().equals(""))
            return;

        activity.getSearchBarET().setText(domain);
        String url = Constants.URL_DOMAIN+domain;

        Thread t = new Thread(
                () -> {
                    util.GETrequest(Constants.DOMAIN_CALLBACK, url);
                }
        );
        t.start();
        try {
            t.join();
        } catch (InterruptedException e) {
            Toast.makeText(activity,Constants.GENERAL_ERROR,Toast.LENGTH_LONG).show();
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
            case Constants.DOMAIN_CALLBACK:
                Gson gson = new Gson();
                Domain domainResponseTMP = null;
                try {
                    domainResponseTMP = gson.fromJson(response, Domain.class);
                } catch (Exception e) {
                    activity.runOnUiThread(
                            () -> {
                                Toast.makeText(activity,Constants.GENERAL_ERROR+"\n"+response,Toast.LENGTH_LONG).show();
                                activity.onBackPressed();
                            }
                    );
                    return;
                }
                Domain domainResponse = domainResponseTMP;
                activity.runOnUiThread(
                        () -> {
                            Picasso.get().load(domainResponse.getLogo()).into(activity.getLogoIV());
                            activity.getSslGradeTV().setText("Grado SSL: "+domainResponse.getSsl_grade());
                            activity.getPreviousSslGradeTV().setText("Grado SSL Previo: "+domainResponse.getPrevious_ssl_grade());
                            activity.getServersChangedTV().setText("Los servidores han cambiado: "+domainResponse.getServers_changed());
                            activity.getIsDownTV().setText("Está caído el servidor: "+domainResponse.getIs_down());
                            List<Server> servers = activity.getServers();
                            servers.clear();
                            for (Server s:domainResponse.getServers()) {
                                servers.add(s);
                            }

                            activity.getAdapter().notifyDataSetChanged();
                        }
                );
                break;
        }
    }

}
