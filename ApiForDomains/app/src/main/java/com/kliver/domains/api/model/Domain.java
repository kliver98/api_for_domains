package com.kliver.domains.api.model;

public class Domain {
    private Server []servers;
    private String serversChanged;
    private String sslGrade;
    private String previousSslGrade;
    private String logo;
    private String title;
    private String isDown;

    public Domain() {}

    public Domain(Server[] servers, String serversChanged, String sslGrade, String previousSslGrade, String logo, String title, String isDown) {
        this.servers = servers;
        this.serversChanged = serversChanged;
        this.sslGrade = sslGrade;
        this.previousSslGrade = previousSslGrade;
        this.logo = logo;
        this.title = title;
        this.isDown = isDown;
    }

    public Server[] getServers() {
        return servers;
    }

    public void setServers(Server[] servers) {
        this.servers = servers;
    }

    public String getServersChanged() {
        return serversChanged;
    }

    public void setServersChanged(String serversChanged) {
        this.serversChanged = serversChanged;
    }

    public String getSslGrade() {
        return sslGrade;
    }

    public void setSslGrade(String sslGrade) {
        this.sslGrade = sslGrade;
    }

    public String getPreviousSslGrade() {
        return previousSslGrade;
    }

    public void setPreviousSslGrade(String previousSslGrade) {
        this.previousSslGrade = previousSslGrade;
    }

    public String getLogo() {
        return logo;
    }

    public void setLogo(String logo) {
        this.logo = logo;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getIsDown() {
        return isDown;
    }

    public void setIsDown(String isDown) {
        this.isDown = isDown;
    }
}
