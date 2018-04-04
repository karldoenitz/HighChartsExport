# -*- coding: utf-8 -*-

import requests
import json


def test():
    request_data = {
        "chart": {
            "polar": True,
            "type": "line"
        },
        "title": {
            "text": "Budget vs spending",
            "x": -80,
            "style": {
                "color": "#000000",
                "fontSize": "18px"
            }
        },
        "pane": {
            "size": "80%"
        },
        "xAxis": {
            "categories": ["测试", "Marketing", "Development", "Customer Support", "Information Technology", "Administration"],
            "tickmarkPlacement": "on",
            "lineWidth": 0,
            "gridLineColor": "#707073",
            "labels": {
                "style": {
                    "color": "#000000"
                }
            },
            "lineColor": "#707073",
            "minorGridLineColor": "#707073",
            "tickColor": "#707073",
        },
        "yAxis": {
            "gridLineColor": "#707073",
            "gridLineInterpolation": "polygon",
            "lineWidth": 0,
            "min": 0,
            "tickInterval": 0.4,
            "tickLength": 1.2
        },
        "tooltip": {
            "shared": True,
            "pointFormat": "<span style=\"color:{series.color}\">{series.name}: <b>${point.y:,.0f}</b><br/>"
        },
        "legend": {
            "align": "right",
            "verticalAlign": "top",
            "y": 70,
            "layout": "vertical"
        },
        "series": [{
            "name": "Allocated Budget",
            "data": [0.60, 0.70, 0.25, 0.75, 0.10, 1.00],
            "pointPlacement": "on",
            "color": "#4a7ebb"
        }, {
            "name": "Actual Spending",
            "data": [0.30, 0.50, 0.40, 0.36, 0.20, 0.10],
            "pointPlacement": "on",
            "color": "#be4b48"
        }]
    }
    headers = {
        "Content-Type": "application/json"
    }
    response = requests.post(
        "http://localhost:8908/export?name=first",
        json.dumps(request_data, ensure_ascii=False),
        headers=headers
    )
    print(response.content)


if __name__ == '__main__':
    test()
