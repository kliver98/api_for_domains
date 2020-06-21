package com.kliver.domains.api.view;

import androidx.appcompat.app.AppCompatActivity;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;
import com.kliver.domains.api.R;
import com.kliver.domains.api.control.HistoryController;
import com.kliver.domains.api.model.adapter.HistoryAdapter;
import com.kliver.domains.api.util.Constants;

import android.content.Intent;
import android.os.Bundle;
import android.widget.ImageView;

import java.util.ArrayList;

public class HistoryActivity extends AppCompatActivity {

    private HistoryController controller;
    private ImageView backArrowIV;
    private RecyclerView itemsRV;
    private HistoryAdapter adapter;
    private ArrayList<String> items;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_history);
        backArrowIV = findViewById(R.id.backArrowIV);
        itemsRV = findViewById(R.id.itemsRV);
        itemsRV.setLayoutManager(new LinearLayoutManager(this));
        itemsRV.setHasFixedSize(true);
        items = new ArrayList<String>();
        adapter = new HistoryAdapter(items);
        adapter.setOnItemClickListener(new HistoryAdapter.OnItemClickListener() {
            @Override
            public void onItemClick(int position) {
                openDomainActivity(items.get(position));
            }
        });
        itemsRV.setAdapter(adapter);
        controller = new HistoryController(this);
    }

    public ImageView getBackArrowIV() {
        return backArrowIV;
    }

    public HistoryAdapter getAdapter() {
        return adapter;
    }

    public ArrayList<String> getItems() {
        return items;
    }

    public void openDomainActivity(String domain) {
        Intent i = new Intent(this,DomainActivity.class);
        i.putExtra(Constants.DOMAIN_VAR_INTENT,domain);
        this.startActivity(i);
    }

    @Override
    public void onBackPressed() {
        super.onBackPressed();
        this.finish();
    }

}
