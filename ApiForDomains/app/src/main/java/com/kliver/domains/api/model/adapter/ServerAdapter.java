package com.kliver.domains.api.model.adapter;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.kliver.domains.api.R;
import com.kliver.domains.api.model.Server;

import java.util.ArrayList;

public class ServerAdapter extends RecyclerView.Adapter<ServerAdapter.ViewHolderServer> {

    private ArrayList<Server> servers;

    public ServerAdapter(ArrayList<Server> servers) {
        this.servers = servers;
    }

    @NonNull
    @Override
    public ViewHolderServer onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.item_server,parent,false);
        return new ViewHolderServer(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolderServer holder, int position) {
        holder.getAddressTV().setText("IP: "+servers.get(position).getAddress());
        holder.getSslGradeTV().setText("Grado SSL: "+servers.get(position).getSsl_grade());
        holder.getCountryTV().setText("País: "+servers.get(position).getCountry());
        holder.getOwnerTV().setText("Dueño: "+servers.get(position).getOwner());
    }

    @Override
    public int getItemCount() {
        return servers.size();
    }

    public class ViewHolderServer extends RecyclerView.ViewHolder {

        private TextView addressTV;
        private TextView sslGradeTV;
        private TextView countryTV;
        private TextView ownerTV;

        public ViewHolderServer(@NonNull View itemView) {
            super(itemView);
            addressTV  = itemView.findViewById(R.id.addressTV);
            sslGradeTV  = itemView.findViewById(R.id.sslGradeTV);
            countryTV  = itemView.findViewById(R.id.countryTV);
            ownerTV  = itemView.findViewById(R.id.ownerTV);
        }

        public TextView getAddressTV() {
            return addressTV;
        }

        public TextView getSslGradeTV() {
            return sslGradeTV;
        }

        public TextView getCountryTV() {
            return countryTV;
        }

        public TextView getOwnerTV() {
            return ownerTV;
        }
    }
}
