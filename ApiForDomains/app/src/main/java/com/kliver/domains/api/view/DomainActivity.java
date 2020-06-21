package com.kliver.domains.api.view;

import androidx.appcompat.app.AppCompatActivity;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.kliver.domains.api.R;
import com.kliver.domains.api.control.DomainController;
import com.kliver.domains.api.model.Server;
import com.kliver.domains.api.model.adapter.ServerAdapter;

import android.os.Bundle;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;

import java.util.ArrayList;

public class DomainActivity extends AppCompatActivity {

    private DomainController controller;
    private ImageView backArrowIV;
    private ImageView searchImgIV;
    private EditText searchBarET;
    private ImageView logoIV;
    private TextView sslGradeTV;
    private TextView previousSslGradeTV;
    private TextView serversChangedTV;
    private TextView isDownTV;
    private RecyclerView serversRV;
    private ArrayList<Server> servers;
    private ServerAdapter adaptader;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_domain);
        backArrowIV = findViewById(R.id.backArrowIV);
        searchImgIV = findViewById(R.id.searchImgIV);
        logoIV = findViewById(R.id.logoIV);
        sslGradeTV = findViewById(R.id.sslGradeTV);
        previousSslGradeTV = findViewById(R.id.previousSslGradeTV);
        serversChangedTV = findViewById(R.id.serversChangedTV);
        isDownTV = findViewById(R.id.isDownTV);
        serversRV = findViewById(R.id.serversRV);
        searchBarET = findViewById(R.id.searchBarET);

        serversRV.setLayoutManager(new LinearLayoutManager(this));
        servers = new ArrayList<Server>();
        adaptader = new ServerAdapter(servers);
        serversRV.setAdapter(adaptader);

        controller = new DomainController(this);
    }

    public ImageView getBackArrowIV() {
        return backArrowIV;
    }

    public ArrayList<Server> getServers() {
        return servers;
    }

    public ServerAdapter getAdapter() {
        return adaptader;
    }

    public ImageView getSearchImgIV() {
        return searchImgIV;
    }

    public EditText getSearchBarET() {
        return searchBarET;
    }

    public ImageView getLogoIV() {
        return logoIV;
    }

    public TextView getSslGradeTV() {
        return sslGradeTV;
    }

    public TextView getPreviousSslGradeTV() {
        return previousSslGradeTV;
    }

    public TextView getServersChangedTV() {
        return serversChangedTV;
    }

    public TextView getIsDownTV() {
        return isDownTV;
    }

    @Override
    public void onBackPressed() {
        super.onBackPressed();
        this.finish();
    }

}
