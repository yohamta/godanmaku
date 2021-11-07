//
//  MobileEbitenViewControllerWithErrorHandling.swift
//  godanmaku
//
//  Created by Yota Hamada on 2020/08/20.
//  Copyright © 2020 Yota Hamada. All rights reserved.
//

import Foundation   
import Mobile

class MobileEbitenViewControllerWithErrorHandling: MobileEbitenViewController {
    override func onError(onGameUpdate err: Error!) {
        print(err ?? "OnGameUpdate Error")
    }
}
